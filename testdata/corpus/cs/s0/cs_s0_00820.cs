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

    public static int Guard7690(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 11;
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
    public delegate void Notifier7691(int value);

    public sealed class Source7691
    {
        public event Notifier7691 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate7692(IEnumerable<int> source)
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
    public interface IHandler7693
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7693(string Label, int Payload);

    public sealed class Handler7693 : IHandler7693
    {
        public void Handle(int value) { }
    }
    public static int Scale7694<T>(T value) where T : struct
    {
        int number = 39;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 68; }
        return number;
    }
}
