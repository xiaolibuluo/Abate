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

    public static int Guard7160(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 98;
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
    public delegate void Notifier7161(int value);

    public sealed class Source7161
    {
        public event Notifier7161 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale7162<T>(T value) where T : struct
    {
        int number = 15;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 41; }
        return number;
    }
    public static int Scale7163<T>(T value) where T : struct
    {
        int number = 22;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public async Task<int> Aggregate7164(IEnumerable<int> source)
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
