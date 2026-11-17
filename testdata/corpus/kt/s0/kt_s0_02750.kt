package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step22700")
fun step22700(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
class Counter22701 {
    var current: Int by Delegates.observable(34) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
@JvmName("step22702")
fun step22702(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
class Counter22703 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
fun <T, R> List<T>.mapChecked22704(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step22705")
fun step22705(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 62
}
