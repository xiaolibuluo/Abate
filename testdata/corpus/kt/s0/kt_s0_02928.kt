package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair24150<A, B>(val first: A, val second: B)

fun combine24150(left: Int, right: Int): Pair24150<Int, Int> {
    return Pair24150(left + 69, right + 17)
}
object Registry24151 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24151 = "registry-76"
    }
}
fun <T> List<T>.secondOrNull24152(): T? = if (size > 1) this[1] else null

fun render24152(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24153<A, B>(val first: A, val second: B)

fun combine24153(left: Int, right: Int): Pair24153<Int, Int> {
    return Pair24153(left + 90, right + 50)
}
object Registry24154 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24154 = "registry-97"
    }
}
class Counter24155 {
    var current: Int by Delegates.observable(7) { _, _, new -> println(new) }

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
