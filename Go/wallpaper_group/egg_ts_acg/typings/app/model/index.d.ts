// This file is created by egg-ts-helper@1.25.9
// Do not modify this file!!!!!!!!!

import 'egg';
import ExportAcg from '../../../app/model/acg';

declare module 'egg' {
  interface IModel {
    Acg: ReturnType<typeof ExportAcg>;
  }
}
