package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull16680(): T? = if (size > 1) this[1] else null

fun render16680(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step16681")
fun step16681(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
object Registry16682 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16682 = "registry-72"
    }
}
class Counter16683 {
    var current: Int by Delegates.observable(79) { _, _, new -> println(new) }

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
