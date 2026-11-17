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

    public interface IHandler23350
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23350(string Label, int Payload);

    public sealed class Handler23350 : IHandler23350
    {
        public void Handle(int value) { }
    }
    public struct Point23351
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 6;
    }

    public static class Grid23351
    {
        public static Point23351 Make(int left, int right)
        {
            return new Point23351 { X = left, Y = right };
        }
    }
    public delegate void Notifier23352(int value);

    public sealed class Source23352
    {
        public event Notifier23352 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard23353(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 20;
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
    public async Task<int> Aggregate23354(IEnumerable<int> source)
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
    public struct Point23355
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 34;
    }

    public static class Grid23355
    {
        public static Point23355 Make(int left, int right)
        {
            return new Point23355 { X = left, Y = right };
        }
    }
}
