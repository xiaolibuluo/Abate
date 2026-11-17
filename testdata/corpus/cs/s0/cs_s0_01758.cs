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

    public static int Guard16380(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 56;
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
    public async Task<int> Aggregate16381(IEnumerable<int> source)
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
    public interface IHandler16382
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16382(string Label, int Payload);

    public sealed class Handler16382 : IHandler16382
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier16383(int value);

    public sealed class Source16383
    {
        public event Notifier16383 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
