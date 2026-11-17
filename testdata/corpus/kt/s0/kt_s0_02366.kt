package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair19440<A, B>(val first: A, val second: B)

fun combine19440(left: Int, right: Int): Pair19440<Int, Int> {
    return Pair19440(left + 57, right + 14)
}
fun <T> List<T>.secondOrNull19441(): T? = if (size > 1) this[1] else null

fun render19441(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step19442")
fun step19442(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 71
}
@JvmName("step19443")
fun step19443(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 78
}
data class Pair19444<A, B>(val first: A, val second: B)

fun combine19444(left: Int, right: Int): Pair19444<Int, Int> {
    return Pair19444(left + 85, right + 58)
}
class Counter19445 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
data class Pair19446<A, B>(val first: A, val second: B)

fun combine19446(left: Int, right: Int): Pair19446<Int, Int> {
    return Pair19446(left + 99, right + 19)
}
class Counter19447 {
    var current: Int by Delegates.observable(9) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull19448(): T? = if (size > 1) this[1] else null

fun render19448(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
