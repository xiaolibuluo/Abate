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

    public interface IHandler4200
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4200(string Label, int Payload);

    public sealed class Handler4200 : IHandler4200
    {
        public void Handle(int value) { }
    }
    public static int Scale4201<T>(T value) where T : struct
    {
        int number = 40;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 32; }
        return number;
    }
    public async Task<int> Aggregate4202(IEnumerable<int> source)
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
    public static int Guard4203(int value)
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
}
