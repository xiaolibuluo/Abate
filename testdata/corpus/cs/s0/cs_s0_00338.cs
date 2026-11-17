using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public async Task<int> Aggregate3200(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public delegate void Notifier3201(int value);

    public sealed class Source3201
    {
        public event Notifier3201 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler3202
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3202(string Label, int Payload);

    public sealed class Handler3202 : IHandler3202
    {
        public void Handle(int value) { }
    }
    public static int Guard3203(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 15;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
}
