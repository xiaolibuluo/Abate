package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step7180")
fun step7180(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
@JvmName("step7181")
fun step7181(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 34
}
object Registry7182 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7182 = "registry-41"
    }
}
object Registry7183 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7183 = "registry-48"
    }
}
class Counter7184 {
    var current: Int by Delegates.observable(55) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull7185(): T? = if (size > 1) this[1] else null

fun render7185(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair7186<A, B>(val first: A, val second: B)

fun combine7186(left: Int, right: Int): Pair7186<Int, Int> {
    return Pair7186(left + 69, right + 38)
}
fun <T> List<T>.secondOrNull7187(): T? = if (size > 1) this[1] else null

fun render7187(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
