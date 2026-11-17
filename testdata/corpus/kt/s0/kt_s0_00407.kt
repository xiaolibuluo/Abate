package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3330<A, B>(val first: A, val second: B)

fun combine3330(left: Int, right: Int): Pair3330<Int, Int> {
    return Pair3330(left + 66, right + 22)
}
data class Pair3331<A, B>(val first: A, val second: B)

fun combine3331(left: Int, right: Int): Pair3331<Int, Int> {
    return Pair3331(left + 73, right + 33)
}
fun <T, R> List<T>.mapChecked3332(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step3333")
fun step3333(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 87
}
object Registry3334 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3334 = "registry-94"
    }
}
fun <T> List<T>.secondOrNull3335(): T? = if (size > 1) this[1] else null

fun render3335(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step3336")
fun step3336(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
class Counter3337 {
    var current: Int by Delegates.observable(18) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull3338(): T? = if (size > 1) this[1] else null

fun render3338(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
