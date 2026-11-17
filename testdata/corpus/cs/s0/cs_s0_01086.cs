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

    public delegate void Notifier10180(int value);

    public sealed class Source10180
    {
        public event Notifier10180 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler10181
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10181(string Label, int Payload);

    public sealed class Handler10181 : IHandler10181
    {
        public void Handle(int value) { }
    }
    public static int Scale10182<T>(T value) where T : struct
    {
        int number = 61;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 20; }
        return number;
    }
    public static int Guard10183(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 68;
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
    public async Task<int> Aggregate10184(IEnumerable<int> source)
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
    public delegate void Notifier10185(int value);

    public sealed class Source10185
    {
        public event Notifier10185 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
