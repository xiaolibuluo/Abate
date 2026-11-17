package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull1190(): T? = if (size > 1) this[1] else null

fun render1190(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry1191 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1191 = "registry-7"
    }
}
data class Pair1192<A, B>(val first: A, val second: B)

fun combine1192(left: Int, right: Int): Pair1192<Int, Int> {
    return Pair1192(left + 14, right + 10)
}
fun <T, R> List<T>.mapChecked1193(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter1194 {
    var current: Int by Delegates.observable(28) { _, _, new -> println(new) }

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
