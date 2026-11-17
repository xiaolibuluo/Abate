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

    public interface IHandler20880
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20880(string Label, int Payload);

    public sealed class Handler20880 : IHandler20880
    {
        public void Handle(int value) { }
    }
    public static int Scale20881<T>(T value) where T : struct
    {
        int number = 93;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 57; }
        return number;
    }
    public static int Guard20882(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 100;
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
    public interface IHandler20883
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20883(string Label, int Payload);

    public sealed class Handler20883 : IHandler20883
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate20884(IEnumerable<int> source)
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
    public struct Point20885
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 24;
    }

    public static class Grid20885
    {
        public static Point20885 Make(int left, int right)
        {
            return new Point20885 { X = left, Y = right };
        }
    }
}
