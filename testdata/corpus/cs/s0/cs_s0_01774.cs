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

    public async Task<int> Aggregate16540(IEnumerable<int> source)
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
    public static int Guard16541(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 77;
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
    public interface IHandler16542
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16542(string Label, int Payload);

    public sealed class Handler16542 : IHandler16542
    {
        public void Handle(int value) { }
    }
    public struct Point16543
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 91;
    }

    public static class Grid16543
    {
        public static Point16543 Make(int left, int right)
        {
            return new Point16543 { X = left, Y = right };
        }
    }
    public struct Point16544
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid16544
    {
        public static Point16544 Make(int left, int right)
        {
            return new Point16544 { X = left, Y = right };
        }
    }
    public struct Point16545
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 8;
    }

    public static class Grid16545
    {
        public static Point16545 Make(int left, int right)
        {
            return new Point16545 { X = left, Y = right };
        }
    }
}
