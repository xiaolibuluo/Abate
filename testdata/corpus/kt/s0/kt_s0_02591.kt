package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull21340(): T? = if (size > 1) this[1] else null

fun render21340(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull21341(): T? = if (size > 1) this[1] else null

fun render21341(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry21342 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21342 = "registry-19"
    }
}
class Counter21343 {
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
data class Pair21344<A, B>(val first: A, val second: B)

fun combine21344(left: Int, right: Int): Pair21344<Int, Int> {
    return Pair21344(left + 33, right + 46)
}
data class Pair21345<A, B>(val first: A, val second: B)

fun combine21345(left: Int, right: Int): Pair21345<Int, Int> {
    return Pair21345(left + 40, right + 57)
}
class Counter21346 {
    var current: Int by Delegates.observable(47) { _, _, new -> println(new) }

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
object Registry21347 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21347 = "registry-54"
    }
}
