package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry19400 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19400 = "registry-5"
    }
}
fun <T> List<T>.secondOrNull19401(): T? = if (size > 1) this[1] else null

fun render19401(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull19402(): T? = if (size > 1) this[1] else null

fun render19402(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter19403 {
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
class Counter19404 {
    var current: Int by Delegates.observable(33) { _, _, new -> println(new) }

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
data class Pair19405<A, B>(val first: A, val second: B)

fun combine19405(left: Int, right: Int): Pair19405<Int, Int> {
    return Pair19405(left + 40, right + 41)
}
data class Pair19406<A, B>(val first: A, val second: B)

fun combine19406(left: Int, right: Int): Pair19406<Int, Int> {
    return Pair19406(left + 47, right + 52)
}
object Registry19407 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19407 = "registry-54"
    }
}
