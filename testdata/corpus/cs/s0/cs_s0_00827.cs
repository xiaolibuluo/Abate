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

    public struct Point7760
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 5;
    }

    public static class Grid7760
    {
        public static Point7760 Make(int left, int right)
        {
            return new Point7760 { X = left, Y = right };
        }
    }
    public interface IHandler7761
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7761(string Label, int Payload);

    public sealed class Handler7761 : IHandler7761
    {
        public void Handle(int value) { }
    }
    public static int Guard7762(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 19;
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
    public async Task<int> Aggregate7763(IEnumerable<int> source)
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
    public struct Point7764
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 33;
    }

    public static class Grid7764
    {
        public static Point7764 Make(int left, int right)
        {
            return new Point7764 { X = left, Y = right };
        }
    }
}
