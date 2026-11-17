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

    public async Task<int> Aggregate11510(IEnumerable<int> source)
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
    public delegate void Notifier11511(int value);

    public sealed class Source11511
    {
        public event Notifier11511 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier11512(int value);

    public sealed class Source11512
    {
        public event Notifier11512 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale11513<T>(T value) where T : struct
    {
        int number = 51;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public delegate void Notifier11514(int value);

    public sealed class Source11514
    {
        public event Notifier11514 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
