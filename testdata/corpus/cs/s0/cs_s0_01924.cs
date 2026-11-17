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

    public struct Point17970
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 86;
    }

    public static class Grid17970
    {
        public static Point17970 Make(int left, int right)
        {
            return new Point17970 { X = left, Y = right };
        }
    }
    public interface IHandler17971
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17971(string Label, int Payload);

    public sealed class Handler17971 : IHandler17971
    {
        public void Handle(int value) { }
    }
    public struct Point17972
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 100;
    }

    public static class Grid17972
    {
        public static Point17972 Make(int left, int right)
        {
            return new Point17972 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate17973(IEnumerable<int> source)
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
    public static int Scale17974<T>(T value) where T : struct
    {
        int number = 17;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 66; }
        return number;
    }
    public static int Guard17975(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 24;
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
