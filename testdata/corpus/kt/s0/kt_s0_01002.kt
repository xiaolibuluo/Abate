package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry8210 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8210 = "registry-8"
    }
}
fun <T, R> List<T>.mapChecked8211(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair8212<A, B>(val first: A, val second: B)

fun combine8212(left: Int, right: Int): Pair8212<Int, Int> {
    return Pair8212(left + 22, right + 44)
}
object Registry8213 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8213 = "registry-29"
    }
}
class Counter8214 {
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
data class Pair8215<A, B>(val first: A, val second: B)

fun combine8215(left: Int, right: Int): Pair8215<Int, Int> {
    return Pair8215(left + 43, right + 16)
}
fun <T> List<T>.secondOrNull8216(): T? = if (size > 1) this[1] else null

fun render8216(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter8217 {
    var current: Int by Delegates.observable(57) { _, _, new -> println(new) }

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
