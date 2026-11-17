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

    public async Task<int> Aggregate11060(IEnumerable<int> source)
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
    public static int Scale11061<T>(T value) where T : struct
    {
        int number = 34;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 15; }
        return number;
    }
    public static int Scale11062<T>(T value) where T : struct
    {
        int number = 41;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 26; }
        return number;
    }
    public struct Point11063
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 48;
    }

    public static class Grid11063
    {
        public static Point11063 Make(int left, int right)
        {
            return new Point11063 { X = left, Y = right };
        }
    }
    public interface IHandler11064
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11064(string Label, int Payload);

    public sealed class Handler11064 : IHandler11064
    {
        public void Handle(int value) { }
    }
}
