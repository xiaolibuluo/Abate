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

    public static int Scale4210<T>(T value) where T : struct
    {
        int number = 46;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 28; }
        return number;
    }
    public static int Scale4211<T>(T value) where T : struct
    {
        int number = 53;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
    public static int Scale4212<T>(T value) where T : struct
    {
        int number = 60;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 50; }
        return number;
    }
    public async Task<int> Aggregate4213(IEnumerable<int> source)
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
    public delegate void Notifier4214(int value);

    public sealed class Source4214
    {
        public event Notifier4214 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
