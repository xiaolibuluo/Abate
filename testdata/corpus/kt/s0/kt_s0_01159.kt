package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9540<A, B>(val first: A, val second: B)

fun combine9540(left: Int, right: Int): Pair9540<Int, Int> {
    return Pair9540(left + 88, right + 38)
}
data class Pair9541<A, B>(val first: A, val second: B)

fun combine9541(left: Int, right: Int): Pair9541<Int, Int> {
    return Pair9541(left + 95, right + 49)
}
class Counter9542 {
    var current: Int by Delegates.observable(5) { _, _, new -> println(new) }

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
data class Pair9543<A, B>(val first: A, val second: B)

fun combine9543(left: Int, right: Int): Pair9543<Int, Int> {
    return Pair9543(left + 12, right + 10)
}
fun <T> List<T>.secondOrNull9544(): T? = if (size > 1) this[1] else null

fun render9544(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter9545 {
    var current: Int by Delegates.observable(26) { _, _, new -> println(new) }

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
