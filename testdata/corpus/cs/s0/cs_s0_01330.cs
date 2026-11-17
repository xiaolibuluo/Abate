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

    public delegate void Notifier12460(int value);

    public sealed class Source12460
    {
        public event Notifier12460 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale12461<T>(T value) where T : struct
    {
        int number = 11;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 19; }
        return number;
    }
    public async Task<int> Aggregate12462(IEnumerable<int> source)
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
    public delegate void Notifier12463(int value);

    public sealed class Source12463
    {
        public event Notifier12463 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
