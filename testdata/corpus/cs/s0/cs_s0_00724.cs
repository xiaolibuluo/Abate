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

    public interface IHandler6790
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6790(string Label, int Payload);

    public sealed class Handler6790 : IHandler6790
    {
        public void Handle(int value) { }
    }
    public static int Scale6791<T>(T value) where T : struct
    {
        int number = 12;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 15; }
        return number;
    }
    public static int Scale6792<T>(T value) where T : struct
    {
        int number = 19;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 26; }
        return number;
    }
    public async Task<int> Aggregate6793(IEnumerable<int> source)
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
    public struct Point6794
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 33;
    }

    public static class Grid6794
    {
        public static Point6794 Make(int left, int right)
        {
            return new Point6794 { X = left, Y = right };
        }
    }
    public delegate void Notifier6795(int value);

    public sealed class Source6795
    {
        public event Notifier6795 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard6796(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 47;
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
