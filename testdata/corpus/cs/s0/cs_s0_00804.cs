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

    public struct Point7540
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 10;
    }

    public static class Grid7540
    {
        public static Point7540 Make(int left, int right)
        {
            return new Point7540 { X = left, Y = right };
        }
    }
    public delegate void Notifier7541(int value);

    public sealed class Source7541
    {
        public event Notifier7541 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate7542(IEnumerable<int> source)
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
