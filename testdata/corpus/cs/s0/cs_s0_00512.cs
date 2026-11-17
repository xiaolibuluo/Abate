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

    public delegate void Notifier4800(int value);

    public sealed class Source4800
    {
        public event Notifier4800 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier4801(int value);

    public sealed class Source4801
    {
        public event Notifier4801 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale4802<T>(T value) where T : struct
    {
        int number = 51;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 36; }
        return number;
    }
    public static int Scale4803<T>(T value) where T : struct
    {
        int number = 58;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public async Task<int> Aggregate4804(IEnumerable<int> source)
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
    public delegate void Notifier4805(int value);

    public sealed class Source4805
    {
        public event Notifier4805 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard4806(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 79;
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
