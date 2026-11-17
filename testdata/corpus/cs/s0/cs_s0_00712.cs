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

    public async Task<int> Aggregate6670(IEnumerable<int> source)
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
    public interface IHandler6671
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6671(string Label, int Payload);

    public sealed class Handler6671 : IHandler6671
    {
        public void Handle(int value) { }
    }
    public static int Scale6672<T>(T value) where T : struct
    {
        int number = 57;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 64; }
        return number;
    }
    public static int Guard6673(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 64;
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
