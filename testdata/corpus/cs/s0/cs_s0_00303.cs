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

    public static int Guard2860(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 37;
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
    public delegate void Notifier2861(int value);

    public sealed class Source2861
    {
        public event Notifier2861 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier2862(int value);

    public sealed class Source2862
    {
        public event Notifier2862 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier2863(int value);

    public sealed class Source2863
    {
        public event Notifier2863 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier2864(int value);

    public sealed class Source2864
    {
        public event Notifier2864 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler2865
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2865(string Label, int Payload);

    public sealed class Handler2865 : IHandler2865
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate2866(IEnumerable<int> source)
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
}
