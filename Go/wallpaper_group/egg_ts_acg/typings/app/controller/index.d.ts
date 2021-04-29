// This file is created by egg-ts-helper@1.25.9
// Do not modify this file!!!!!!!!!

import 'egg';
import ExportHome from '../../../app/controller/home';
import ExportList from '../../../app/controller/list';
import ExportRandom from '../../../app/controller/random';
import ExportTag from '../../../app/controller/tag';

declare module 'egg' {
  interface IController {
    home: ExportHome;
    list: ExportList;
    random: ExportRandom;
    tag: ExportTag;
  }
}
