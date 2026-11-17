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

    public async Task<int> Aggregate2340(IEnumerable<int> source)
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
    public delegate void Notifier2341(int value);

    public sealed class Source2341
    {
        public event Notifier2341 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard2342(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 54;
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
    public interface IHandler2343
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2343(string Label, int Payload);

    public sealed class Handler2343 : IHandler2343
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier2344(int value);

    public sealed class Source2344
    {
        public event Notifier2344 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
