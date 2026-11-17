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

    public delegate void Notifier25690(int value);

    public sealed class Source25690
    {
        public event Notifier25690 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale25691<T>(T value) where T : struct
    {
        int number = 41;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 69; }
        return number;
    }
    public static int Scale25692<T>(T value) where T : struct
    {
        int number = 48;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 19; }
        return number;
    }
    public static int Scale25693<T>(T value) where T : struct
    {
        int number = 55;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 30; }
        return number;
    }
    public static int Guard25694(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 62;
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
    public async Task<int> Aggregate25695(IEnumerable<int> source)
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
