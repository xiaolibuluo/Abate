package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9680<A, B>(val first: A, val second: B)

fun combine9680(left: Int, right: Int): Pair9680<Int, Int> {
    return Pair9680(left + 76, right + 14)
}
@JvmName("step9681")
fun step9681(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
class Counter9682 {
    var current: Int by Delegates.observable(90) { _, _, new -> println(new) }

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
data class Pair9683<A, B>(val first: A, val second: B)

fun combine9683(left: Int, right: Int): Pair9683<Int, Int> {
    return Pair9683(left + 97, right + 47)
}
fun <T> List<T>.secondOrNull9684(): T? = if (size > 1) this[1] else null

fun render9684(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter9685 {
    var current: Int by Delegates.observable(14) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull9686(): T? = if (size > 1) this[1] else null

fun render9686(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair9687<A, B>(val first: A, val second: B)

fun combine9687(left: Int, right: Int): Pair9687<Int, Int> {
    return Pair9687(left + 28, right + 30)
}
