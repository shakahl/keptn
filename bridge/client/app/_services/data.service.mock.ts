import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";

import {DataService} from "./data.service";
import {ApiService} from "./api.service";

import {Root} from "../_models/root";
import {Project} from "../_models/project";
import {DateUtil} from "../_utils/date.utils";

import {KeptnInfo} from "./_mockData/keptnInfo.mock";
import {Projects} from "./_mockData/projects.mock";
import {RootEvents} from "./_mockData/roots.mock";
import {Traces} from "./_mockData/traces.mock";
import {Evaluations} from "./_mockData/evaluations.mock";
import {Trace} from "../_models/trace";
import {Sequence} from '../_models/sequence';

@Injectable({
  providedIn: 'root'
})
export class DataServiceMock extends DataService {

  constructor(apiService: ApiService) {
    super(apiService);
  }

  public loadKeptnInfo() {
    this._keptnInfo.next(KeptnInfo);
  }

  public loadProjects() {
    this._projects.next(Projects);
  }

  public loadProject(projectName) {
    this._projects.next([...Projects]);
  }

  public loadRoots(project: Project) {
    project.roots = [...RootEvents || [], ...project.roots || []].sort(DateUtil.compareTraceTimesAsc);
    project.stages.forEach(stage => {
      stage.services.forEach(service => {
        service.roots = project.roots.filter(s => s.getService() == service.serviceName && s.getStages().includes(stage.stageName));
        service.openApprovals = service.roots.reduce((openApprovals, root) => {
          const approval = root.getPendingApproval(stage.stageName);
          if(approval) {
            openApprovals.push(approval);
          }
          return openApprovals;
        }, []);
      });
    });
    this._roots.next(project.roots);
  }

  public loadTraces(sequence: Sequence) {
    sequence.traces = [...Traces || [], ...sequence.traces || []];
    this._sequences.next([...this._sequences.getValue()]);
  }

  public loadTracesByContext(shkeptncontext: string) {
    this._traces.next(Traces.filter(t => t.shkeptncontext === shkeptncontext));
  }

  public loadEvaluationResults(event: Trace) {
    this._evaluationResults.next({
      type: 'evaluationHistory',
      triggerEvent: event,
      traces: Evaluations
    });
  }

}

