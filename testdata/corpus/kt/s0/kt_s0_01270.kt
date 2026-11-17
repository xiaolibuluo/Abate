package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10500<A, B>(val first: A, val second: B)

fun combine10500(left: Int, right: Int): Pair10500<Int, Int> {
    return Pair10500(left + 75, right + 39)
}
object Registry10501 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10501 = "registry-82"
    }
}
object Registry10502 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10502 = "registry-89"
    }
}
fun <T> List<T>.secondOrNull10503(): T? = if (size > 1) this[1] else null

fun render10503(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step10504")
fun step10504(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 6
}
@JvmName("step10505")
fun step10505(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 13
}
fun <T> List<T>.secondOrNull10506(): T? = if (size > 1) this[1] else null

fun render10506(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
