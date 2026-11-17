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

    public delegate void Notifier22290(int value);

    public sealed class Source22290
    {
        public event Notifier22290 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale22291<T>(T value) where T : struct
    {
        int number = 83;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 68; }
        return number;
    }
    public async Task<int> Aggregate22292(IEnumerable<int> source)
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
    public interface IHandler22293
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22293(string Label, int Payload);

    public sealed class Handler22293 : IHandler22293
    {
        public void Handle(int value) { }
    }
    public static int Guard22294(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 7;
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
