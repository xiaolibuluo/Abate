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

    public static int Guard23660(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 14;
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
    public interface IHandler23661
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23661(string Label, int Payload);

    public sealed class Handler23661 : IHandler23661
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate23662(IEnumerable<int> source)
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
    public static int Scale23663<T>(T value) where T : struct
    {
        int number = 35;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 12; }
        return number;
    }
}
