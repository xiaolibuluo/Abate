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

    public interface IHandler22700
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22700(string Label, int Payload);

    public sealed class Handler22700 : IHandler22700
    {
        public void Handle(int value) { }
    }
    public struct Point22701
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 34;
    }

    public static class Grid22701
    {
        public static Point22701 Make(int left, int right)
        {
            return new Point22701 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate22702(IEnumerable<int> source)
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
    public async Task<int> Aggregate22703(IEnumerable<int> source)
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
