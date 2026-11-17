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

    public static int Guard9070(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 59;
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
    public async Task<int> Aggregate9071(IEnumerable<int> source)
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
    public delegate void Notifier9072(int value);

    public sealed class Source9072
    {
        public event Notifier9072 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale9073<T>(T value) where T : struct
    {
        int number = 80;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
}
