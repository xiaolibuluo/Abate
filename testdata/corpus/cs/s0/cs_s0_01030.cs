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

    public delegate void Notifier9650(int value);

    public sealed class Source9650
    {
        public event Notifier9650 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier9651(int value);

    public sealed class Source9651
    {
        public event Notifier9651 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point9652
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 51;
    }

    public static class Grid9652
    {
        public static Point9652 Make(int left, int right)
        {
            return new Point9652 { X = left, Y = right };
        }
    }
    public static int Guard9653(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 58;
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
    public async Task<int> Aggregate9654(IEnumerable<int> source)
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
