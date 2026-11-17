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

    public delegate void Notifier5810(int value);

    public sealed class Source5810
    {
        public event Notifier5810 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale5811<T>(T value) where T : struct
    {
        int number = 96;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 61; }
        return number;
    }
    public static int Scale5812<T>(T value) where T : struct
    {
        int number = 6;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
    public async Task<int> Aggregate5813(IEnumerable<int> source)
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
    public interface IHandler5814
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5814(string Label, int Payload);

    public sealed class Handler5814 : IHandler5814
    {
        public void Handle(int value) { }
    }
    public static int Scale5815<T>(T value) where T : struct
    {
        int number = 27;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
    public static int Guard5816(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 34;
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
