package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair5220<A, B>(val first: A, val second: B)

fun combine5220(left: Int, right: Int): Pair5220<Int, Int> {
    return Pair5220(left + 98, right + 64)
}
data class Pair5221<A, B>(val first: A, val second: B)

fun combine5221(left: Int, right: Int): Pair5221<Int, Int> {
    return Pair5221(left + 8, right + 14)
}
@JvmName("step5222")
fun step5222(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
fun <T> List<T>.secondOrNull5223(): T? = if (size > 1) this[1] else null

fun render5223(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair5224<A, B>(val first: A, val second: B)

fun combine5224(left: Int, right: Int): Pair5224<Int, Int> {
    return Pair5224(left + 29, right + 47)
}
class Counter5225 {
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
fun <T> List<T>.secondOrNull5226(): T? = if (size > 1) this[1] else null

fun render5226(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step5227")
fun step5227(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
