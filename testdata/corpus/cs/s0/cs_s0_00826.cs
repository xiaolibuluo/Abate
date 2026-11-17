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

    public static int Scale7750<T>(T value) where T : struct
    {
        int number = 89;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 66; }
        return number;
    }
    public interface IHandler7751
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7751(string Label, int Payload);

    public sealed class Handler7751 : IHandler7751
    {
        public void Handle(int value) { }
    }
    public struct Point7752
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 6;
    }

    public static class Grid7752
    {
        public static Point7752 Make(int left, int right)
        {
            return new Point7752 { X = left, Y = right };
        }
    }
    public interface IHandler7753
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7753(string Label, int Payload);

    public sealed class Handler7753 : IHandler7753
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate7754(IEnumerable<int> source)
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
    public async Task<int> Aggregate7755(IEnumerable<int> source)
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
