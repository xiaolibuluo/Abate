package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry9510 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9510 = "registry-49"
    }
}
fun <T> List<T>.secondOrNull9511(): T? = if (size > 1) this[1] else null

fun render9511(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry9512 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9512 = "registry-63"
    }
}
fun <T> List<T>.secondOrNull9513(): T? = if (size > 1) this[1] else null

fun render9513(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked9514(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step9515")
fun step9515(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 84
}
