package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16860<A, B>(val first: A, val second: B)

fun combine16860(left: Int, right: Int): Pair16860<Int, Int> {
    return Pair16860(left + 98, right + 38)
}
class Counter16861 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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
data class Pair16862<A, B>(val first: A, val second: B)

fun combine16862(left: Int, right: Int): Pair16862<Int, Int> {
    return Pair16862(left + 15, right + 60)
}
fun <T> List<T>.secondOrNull16863(): T? = if (size > 1) this[1] else null

fun render16863(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter16864 {
    var current: Int by Delegates.observable(29) { _, _, new -> println(new) }

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
class Counter16865 {
    var current: Int by Delegates.observable(36) { _, _, new -> println(new) }

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
