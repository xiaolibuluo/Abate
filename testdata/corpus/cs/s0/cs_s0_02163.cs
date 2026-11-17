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

    public static int Scale20190<T>(T value) where T : struct
    {
        int number = 62;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 51; }
        return number;
    }
    public async Task<int> Aggregate20191(IEnumerable<int> source)
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
    public delegate void Notifier20192(int value);

    public sealed class Source20192
    {
        public event Notifier20192 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier20193(int value);

    public sealed class Source20193
    {
        public event Notifier20193 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
