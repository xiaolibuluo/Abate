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

    public async Task<int> Aggregate2210(IEnumerable<int> source)
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
    public delegate void Notifier2211(int value);

    public sealed class Source2211
    {
        public event Notifier2211 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier2212(int value);

    public sealed class Source2212
    {
        public event Notifier2212 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale2213<T>(T value) where T : struct
    {
        int number = 86;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 64; }
        return number;
    }
}
