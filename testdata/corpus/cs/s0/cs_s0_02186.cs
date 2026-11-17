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

    public static int Guard20410(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 57;
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
    public struct Point20411
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 64;
    }

    public static class Grid20411
    {
        public static Point20411 Make(int left, int right)
        {
            return new Point20411 { X = left, Y = right };
        }
    }
    public struct Point20412
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 71;
    }

    public static class Grid20412
    {
        public static Point20412 Make(int left, int right)
        {
            return new Point20412 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate20413(IEnumerable<int> source)
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
    public delegate void Notifier20414(int value);

    public sealed class Source20414
    {
        public event Notifier20414 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate20415(IEnumerable<int> source)
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
