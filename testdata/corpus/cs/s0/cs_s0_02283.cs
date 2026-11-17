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

    public static int Guard21310(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 63;
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
    public delegate void Notifier21311(int value);

    public sealed class Source21311
    {
        public event Notifier21311 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale21312<T>(T value) where T : struct
    {
        int number = 77;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 64; }
        return number;
    }
    public static int Scale21313<T>(T value) where T : struct
    {
        int number = 84;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 14; }
        return number;
    }
    public async Task<int> Aggregate21314(IEnumerable<int> source)
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
