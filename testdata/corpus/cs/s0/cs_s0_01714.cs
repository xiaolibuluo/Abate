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

    public struct Point16010
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 60;
    }

    public static class Grid16010
    {
        public static Point16010 Make(int left, int right)
        {
            return new Point16010 { X = left, Y = right };
        }
    }
    public delegate void Notifier16011(int value);

    public sealed class Source16011
    {
        public event Notifier16011 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate16012(IEnumerable<int> source)
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
    public interface IHandler16013
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16013(string Label, int Payload);

    public sealed class Handler16013 : IHandler16013
    {
        public void Handle(int value) { }
    }
    public static int Scale16014<T>(T value) where T : struct
    {
        int number = 88;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 36; }
        return number;
    }
    public interface IHandler16015
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16015(string Label, int Payload);

    public sealed class Handler16015 : IHandler16015
    {
        public void Handle(int value) { }
    }
}
