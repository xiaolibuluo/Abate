package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull20660(): T? = if (size > 1) this[1] else null

fun render20660(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter20661 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

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
@JvmName("step20662")
fun step20662(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
fun <T> List<T>.secondOrNull20663(): T? = if (size > 1) this[1] else null

fun render20663(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry20664 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20664 = "registry-22"
    }
}
