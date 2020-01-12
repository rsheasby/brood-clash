/*
 * Brood Clash
 * Backend API for Brood Clash
 *
 * OpenAPI spec version: 0.1.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.10
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.BroodClash);
  }
}(this, function(expect, BroodClash) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new BroodClash.DefaultApi();
  });

  describe('(package)', function() {
    describe('DefaultApi', function() {
      describe('authTest', function() {
        it('should call authTest successfully', function(done) {
          // TODO: uncomment authTest call
          /*

          instance.authTest().then(function(data) {

            done();
          }, function(error) {
            done(error);
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('createQuestions', function() {
        it('should call createQuestions successfully', function(done) {
          // TODO: uncomment, update parameter values for createQuestions call
          /*
          var questions = [new BroodClash.Question()];
          questions[0].text = "";
          questions[0].answers = [new BroodClash.Answer()];
          questions[0].answers[0].text = "";
          questions[0].answers[0].points = 0;
          questions[0].answers[0].revealed = false;

          instance.createQuestions(questions).then(function(data) {

            done();
          }, function(error) {
            done(error);
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('getQuestion', function() {
        it('should call getQuestion successfully', function(done) {
          // TODO: uncomment, update parameter values for getQuestion call and complete the assertions
          /*
          var id = 789;

          instance.getQuestion(id).then(function(data) {
            // TODO: update response assertions
            expect(data).to.be.a(BroodClash.Question);
            expect(data.id).to.be.a('number');
            // expect(data.id).to.be("0");
            expect(data.text).to.be.a('string');
            // expect(data.text).to.be("");
            {
              let dataCtr = data.answers;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(BroodClash.Answer);
                expect(data.id).to.be.a('number');
                // expect(data.id).to.be("0");
                expect(data.text).to.be.a('string');
                // expect(data.text).to.be("");
                expect(data.points).to.be.a('number');
                // expect(data.points).to.be(0);
                expect(data.revealed).to.be.a('boolean');
                // expect(data.revealed).to.be(false);
              }
            }

            done();
          }, function(error) {
            done(error);
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('getQuestions', function() {
        it('should call getQuestions successfully', function(done) {
          // TODO: uncomment getQuestions call and complete the assertions
          /*

          instance.getQuestions().then(function(data) {
            // TODO: update response assertions
            let dataCtr = data;
            expect(dataCtr).to.be.an(Array);
            expect(dataCtr).to.not.be.empty();
            for (let p in dataCtr) {
              let data = dataCtr[p];
              expect(data).to.be.a(BroodClash.Question);
              expect(data.id).to.be.a('number');
              // expect(data.id).to.be("0");
              expect(data.text).to.be.a('string');
              // expect(data.text).to.be("");
              {
                let dataCtr = data.answers;
                expect(dataCtr).to.be.an(Array);
                expect(dataCtr).to.not.be.empty();
                for (let p in dataCtr) {
                  let data = dataCtr[p];
                  expect(data).to.be.a(BroodClash.Answer);
                  expect(data.id).to.be.a('number');
                  // expect(data.id).to.be("0");
                  expect(data.text).to.be.a('string');
                  // expect(data.text).to.be("");
                  expect(data.points).to.be.a('number');
                  // expect(data.points).to.be(0);
                  expect(data.revealed).to.be.a('boolean');
                  // expect(data.revealed).to.be(false);
                }
              }
            }

            done();
          }, function(error) {
            done(error);
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('revealAnswer', function() {
        it('should call revealAnswer successfully', function(done) {
          // TODO: uncomment, update parameter values for revealAnswer call
          /*
          var questionId = 789;
          var answerId = 789;

          instance.revealAnswer(questionId, answerId).then(function(data) {

            done();
          }, function(error) {
            done(error);
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('websocket', function() {
        it('should call websocket successfully', function(done) {
          // TODO: uncomment websocket call
          /*

          instance.websocket().then(function(data) {

            done();
          }, function(error) {
            done(error);
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
    });
  });

}));
