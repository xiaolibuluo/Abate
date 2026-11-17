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

    public struct Point9710
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 18;
    }

    public static class Grid9710
    {
        public static Point9710 Make(int left, int right)
        {
            return new Point9710 { X = left, Y = right };
        }
    }
    public static int Scale9711<T>(T value) where T : struct
    {
        int number = 25;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 46; }
        return number;
    }
    public async Task<int> Aggregate9712(IEnumerable<int> source)
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
    public struct Point9713
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 39;
    }

    public static class Grid9713
    {
        public static Point9713 Make(int left, int right)
        {
            return new Point9713 { X = left, Y = right };
        }
    }
    public interface IHandler9714
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9714(string Label, int Payload);

    public sealed class Handler9714 : IHandler9714
    {
        public void Handle(int value) { }
    }
    public static int Guard9715(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 53;
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
