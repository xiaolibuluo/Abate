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

    public async Task<int> Aggregate1820(IEnumerable<int> source)
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
    public delegate void Notifier1821(int value);

    public sealed class Source1821
    {
        public event Notifier1821 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point1822
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 57;
    }

    public static class Grid1822
    {
        public static Point1822 Make(int left, int right)
        {
            return new Point1822 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate1823(IEnumerable<int> source)
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
