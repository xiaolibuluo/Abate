package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step19860")
fun step19860(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
class Counter19861 {
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
object Registry19862 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19862 = "registry-35"
    }
}
object Registry19863 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19863 = "registry-42"
    }
}
data class Pair19864<A, B>(val first: A, val second: B)

fun combine19864(left: Int, right: Int): Pair19864<Int, Int> {
    return Pair19864(left + 49, right + 47)
}
fun <T> List<T>.secondOrNull19865(): T? = if (size > 1) this[1] else null

fun render19865(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull19866(): T? = if (size > 1) this[1] else null

fun render19866(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry19867 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19867 = "registry-70"
    }
}
