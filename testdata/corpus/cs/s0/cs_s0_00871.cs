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

    public interface IHandler8170
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8170(string Label, int Payload);

    public sealed class Handler8170 : IHandler8170
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier8171(int value);

    public sealed class Source8171
    {
        public event Notifier8171 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale8172<T>(T value) where T : struct
    {
        int number = 67;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 16; }
        return number;
    }
    public static int Guard8173(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 74;
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
    public async Task<int> Aggregate8174(IEnumerable<int> source)
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
