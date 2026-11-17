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

    public interface IHandler25740
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25740(string Label, int Payload);

    public sealed class Handler25740 : IHandler25740
    {
        public void Handle(int value) { }
    }
    public static int Guard25741(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 9;
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
    public interface IHandler25742
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25742(string Label, int Payload);

    public sealed class Handler25742 : IHandler25742
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate25743(IEnumerable<int> source)
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
    public struct Point25744
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 30;
    }

    public static class Grid25744
    {
        public static Point25744 Make(int left, int right)
        {
            return new Point25744 { X = left, Y = right };
        }
    }
}
